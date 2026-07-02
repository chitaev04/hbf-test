import os

output_file = "project_dump.txt"
# Папки и расширения, которые стоит исключить
EXCLUDE_DIRS = {'.git', '__pycache__', 'node_modules', '.venv', 'venv'}
EXCLUDE_EXT  = {'.png', '.jpg', '.jpeg', '.gif', '.ico', '.pdf',
                '.zip', '.exe', '.pyc'}

with open(output_file, "w", encoding="utf-8") as out:
    for root, dirs, files in os.walk("."):
        # Исключаем ненужные директории
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]

        level = root.replace(".", "").count(os.sep)
        indent = "  " * level
        out.write(f"{indent}[DIR] {os.path.basename(root)}/\n")

        for file in files:
            filepath = os.path.join(root, file)
            ext = os.path.splitext(file)[1].lower()
            out.write(f"{indent}  [FILE] {file}\n")

            if ext not in EXCLUDE_EXT:
                try:
                    with open(filepath, "r", encoding="utf-8") as f:
                        content = f.read()
                    out.write(f"{'─'*60}\n")
                    out.write(content)
                    out.write(f"\n{'─'*60}\n")
                except Exception as e:
                    out.write(f"  [не удалось прочитать: {e}]\n")

print(f"Готово! Файл сохранён: {output_file}")