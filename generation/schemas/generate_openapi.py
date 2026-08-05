#!/usr/bin/env python3
"""Builds an OpenAPI yaml file from a template and sibling directories.
Nested directories mirror YAML nesting, with the exception of leaf files with
top level keys with a "/" prefix, which are flattened, ignoring nesting.

Usage:
    build_openapi.py TEMPLATE OUTPUT [--root ROOT_DIR]
"""

import argparse
from pathlib import Path


def indent(text: str, n: int) -> str:
    pad = " " * n
    return "\n".join(pad + line if line.strip() else "" for line in text.splitlines())


def key_of(body: str) -> str:
    first = next((line for line in body.splitlines() if line.strip()), "")
    return first.split(":", 1)[0].strip(" '\"")


def build_tree(top_dir: Path) -> dict:
    tree: dict = {}
    for f in sorted(top_dir.rglob("*.yaml")):
        body = f.read_text().rstrip("\n")
        parts = (
            () if key_of(body).startswith("/") else f.relative_to(top_dir).parts[:-1]
        )
        node = tree
        for part in parts:
            node = node.setdefault(part, {})
        node.setdefault("_files", []).append(body)
    return tree


def render(tree: dict, depth: int) -> list:
    items = [(k, tree[k]) for k in sorted(tree) if k != "_files"]
    items += [(None, body) for body in tree.get("_files", [])]
    lines = []
    for i, (key, val) in enumerate(items):
        if key is None:
            lines.append(indent(val, depth * 2))
        else:
            lines.append(f"{' ' * (depth * 2)}{key}:")
            lines.extend(render(val, depth + 1))
        if i < len(items) - 1:
            lines.append("")
    return lines


def build(template_path: Path, output_path: Path, root: Path) -> None:
    output = template_path.read_text().rstrip() + "\n"
    for top_dir in sorted(
        p for p in root.iterdir() if p.is_dir() and not p.name.startswith(".")
    ):
        lines = [f"{top_dir.name}:"] + render(build_tree(top_dir), 1)
        output += "\n" + "\n".join(lines).rstrip() + "\n"
    output_path.write_text(output)
    print(f"Generated {output_path}")


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "template", type=Path, help="Path to the openapi template yaml file"
    )
    parser.add_argument(
        "output", type=Path, help="Path to write the generated yaml file"
    )
    parser.add_argument(
        "--root",
        type=Path,
        default=None,
        help="Directory containing sibling dirs to merge in (default: template's parent dir)",
    )
    args = parser.parse_args()

    root = args.root if args.root is not None else args.template.parent
    build(args.template, args.output, root)


if __name__ == "__main__":
    main()
