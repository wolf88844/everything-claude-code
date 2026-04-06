const fs = require("fs");
const path = require("path");
const os = require("os");

const repoRoot = path.resolve(__dirname, "..");
const sourceRoot = path.join(repoRoot, "shared-workflows");
const targetRoot = path.join(os.homedir(), ".codex", "skills");

const workflows = [
  "wechat-growth-workflow",
  "social-inspiration-workflow",
  "xiaohongshu-content",
  "personal-ip-strategy",
  "article-writing",
];

function copyDir(src, dest) {
  fs.mkdirSync(dest, { recursive: true });
  for (const entry of fs.readdirSync(src, { withFileTypes: true })) {
    const srcPath = path.join(src, entry.name);
    const destPath = path.join(dest, entry.name);
    if (entry.isDirectory()) {
      copyDir(srcPath, destPath);
    } else {
      fs.copyFileSync(srcPath, destPath);
    }
  }
}

function main() {
  if (!fs.existsSync(sourceRoot)) {
    throw new Error(`Missing shared workflows directory: ${sourceRoot}`);
  }

  fs.mkdirSync(targetRoot, { recursive: true });

  const copied = [];
  for (const workflow of workflows) {
    const src = path.join(sourceRoot, workflow);
    const dest = path.join(targetRoot, workflow);

    if (!fs.existsSync(src)) {
      console.warn(`Skipping missing workflow: ${workflow}`);
      continue;
    }

    copyDir(src, dest);
    copied.push(workflow);
  }

  console.log("Synced shared workflows to Codex skills:");
  for (const workflow of copied) {
    console.log(`- ${workflow}`);
  }
  console.log(`Target: ${targetRoot}`);
}

main();
