import json
import sys
import re
import hashlib
from collections import defaultdict, Counter


def pascal_case(name: str) -> str:
    parts = re.split(r"[^a-zA-Z0-9]", name)
    return "".join(p.capitalize() for p in parts if p)


class TypeStat:

    def __init__(self):
        self.counter = Counter()

    def add(self, t):
        self.counter[t] += 1

    def total(self):
        return sum(self.counter.values())

    def dominant(self):
        if len(self.counter) == 1:
            return list(self.counter.keys())[0]
        return "interface{}"

    def report(self):
        total = self.total()
        parts = []
        for k, v in self.counter.items():
            p = round(v / total * 100, 1)
            parts.append(f"{k}({p}%)")
        return " ".join(parts)


class StructAnalyzer:

    def __init__(self):

        self.struct_defs = {}
        self.struct_cache = {}

        self.field_struct_pool = defaultdict(list)

        self.struct_id = 1

    def struct_hash(self, fields):

        sig = sorted((k, fields[k]["type"]) for k in fields)

        return hashlib.md5(str(sig).encode()).hexdigest()

    def detect_map(self, obj):

        if len(obj) < 2:
            return None

        objs = []

        for v in obj.values():
            if not isinstance(v, dict):
                return None
            objs.append(v)

        merged = self.merge_objects(objs)

        name = self.register_struct(merged, "Item")

        return f"map[string]{name}"

    def analyze_value(self, value, field_name=None):

        if value is None:
            return "interface{}"

        if isinstance(value, str):
            return "string"

        if isinstance(value, bool):
            return "bool"

        if isinstance(value, (int, float)):
            return "float64"

        if isinstance(value, list):
            return self.analyze_array(value, field_name)

        if isinstance(value, dict):

            map_type = self.detect_map(value)

            if map_type:
                return map_type

            if field_name:
                self.field_struct_pool[field_name].append(value)
                return pascal_case(field_name)

            merged = self.merge_objects([value])

            return self.register_struct(merged)

        return "interface{}"

    def analyze_array(self, arr, field_name):

        if not arr:
            return "[]interface{}"

        objs = []
        types = []

        for v in arr:

            if isinstance(v, dict):
                objs.append(v)
            else:
                types.append(self.analyze_value(v))

        if objs:

            if field_name:
                self.field_struct_pool[field_name].extend(objs)
                return f"[]{pascal_case(field_name)}"

            merged = self.merge_objects(objs)

            name = self.register_struct(merged)

            return f"[]{name}"

        stat = TypeStat()

        for t in types:
            stat.add(t)

        t = stat.dominant()

        if t == "interface{}":
            return "[]interface{}"

        return f"[]{t}"

    def merge_objects(self, objs):

        field_stats = defaultdict(TypeStat)

        field_count = Counter()

        for o in objs:

            for k, v in o.items():

                t = self.analyze_value(v, k)

                field_stats[k].add(t)

                field_count[k] += 1

        fields = {}

        total = len(objs)

        for k, stat in field_stats.items():

            optional = field_count[k] < total

            t = stat.dominant()

            conflict = None

            if t == "interface{}":
                conflict = stat.report()

            fields[k] = {"type": t, "optional": optional, "conflict": conflict}

        return fields

    def register_struct(self, fields, field_name=None):

        h = self.struct_hash(fields)

        if h in self.struct_cache:
            return self.struct_cache[h]

        if field_name:
            name = pascal_case(field_name)
        else:
            name = f"AutoStruct{self.struct_id}"

        if name in self.struct_defs:
            name = f"{name}{self.struct_id}"

        self.struct_id += 1

        self.struct_defs[name] = fields

        self.struct_cache[h] = name

        return name

    def finalize_field_structs(self):
        pool_copy = dict(self.field_struct_pool)

        for field, objs in pool_copy.items():

            merged = self.merge_objects(objs)

            name = pascal_case(field)

            if name in self.struct_defs:
                continue

            self.struct_defs[name] = merged

            self.struct_cache[self.struct_hash(merged)] = name

    def extract_deps(self, fields):

        deps = set()

        for v in fields.values():

            t = v["type"]

            t = t.replace("[]", "")

            if t.startswith("map[string]"):
                t = t.replace("map[string]", "")

            if t in self.struct_defs:
                deps.add(t)

        return deps

    def sort_structs(self):

        deps = {k: self.extract_deps(v) for k, v in self.struct_defs.items()}

        visited = set()
        order = []

        def dfs(n):

            if n in visited:
                return

            visited.add(n)

            for d in deps[n]:
                dfs(d)

            order.append(n)

        for n in self.struct_defs:
            dfs(n)

        return order

    def generate_go(self):

        lines = []

        lines.append("package models\n")

        order = self.sort_structs()

        for name in order:

            fields = self.struct_defs[name]

            lines.append(f"type {name} struct {{")

            for k, v in fields.items():

                go_name = pascal_case(k)

                t = v["type"]

                optional = v["optional"]

                if optional and not t.startswith("[]") and not t.startswith("map"):
                    t = "*" + t

                tag = f'`json:"{k}'

                if optional:
                    tag += ",omitempty"

                tag += '"`'

                comment = ""

                if v["conflict"]:
                    comment = f" // conflict: {v['conflict']}"

                lines.append(f"    {go_name} {t} {tag}{comment}")

            lines.append("}\n")

        return "\n".join(lines)


def main(file, key, out):

    with open(file, "r", encoding="utf8") as f:
        data = json.load(f)

    if key not in data:
        print("top key not found")
        return

    root = data[key]

    if not isinstance(root, dict):
        print("top key must be object")
        return

    analyzer = StructAnalyzer()

    analyzer.analyze_value(root, key)

    analyzer.finalize_field_structs()

    go_code = analyzer.generate_go()

    with open(out, "w", encoding="utf8") as f:
        f.write(go_code)

    print("generated:", out)


if __name__ == "__main__":
    main(
        "rawData/data-raw-dump.json", "container", "rawData/container.go"
    )
