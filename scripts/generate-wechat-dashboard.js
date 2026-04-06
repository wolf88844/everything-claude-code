const fs = require('fs');
const path = require('path');

const dashboardDir = path.join(process.cwd(), 'analysis', 'wechat-dashboard');
const sourcePath = path.join(dashboardDir, 'source.json');
const sourcesDir = path.join(dashboardDir, 'sources');
const outputPath = path.join(dashboardDir, 'data.js');

function readJson(filePath) {
  const raw = fs.readFileSync(filePath, 'utf8').replace(/^\uFEFF/, '');
  return JSON.parse(raw);
}

function loadRawData() {
  if (fs.existsSync(sourcesDir)) {
    const metaPath = path.join(sourcesDir, 'meta.json');
    const topicsPath = path.join(sourcesDir, 'topics.json');
    const samplesPath = path.join(sourcesDir, 'samples.json');
    const clusterPath = path.join(sourcesDir, 'cluster.json');

    const meta = fs.existsSync(metaPath) ? readJson(metaPath) : {};
    const topics = fs.existsSync(topicsPath) ? readJson(topicsPath) : {};
    const samples = fs.existsSync(samplesPath) ? readJson(samplesPath) : {};
    const cluster = fs.existsSync(clusterPath) ? readJson(clusterPath) : {};

    return {
      ...meta,
      ...topics,
      ...samples,
      ...cluster
    };
  }

  return readJson(sourcePath);
}

function computeMetrics(data) {
  const topics = Array.isArray(data.topics) ? data.topics : [];
  const draftLikeCount = topics.filter((topic) => ['草稿', '已重写'].includes(topic.status)).length;
  const validatedCount = topics.filter((topic) => topic.status === '已验证').length;

  return [
    { value: topics.length, label: '可管理选题' },
    { value: draftLikeCount, label: '草稿 / 重写' },
    { value: validatedCount, label: '已验证样本' }
  ];
}

function buildDashboardData(rawData) {
  return {
    ...rawData,
    metrics: computeMetrics(rawData)
  };
}

function writeDashboardData(data, filePath) {
  const content = `window.DASHBOARD_DATA = ${JSON.stringify(data, null, 2)};\n`;
  fs.writeFileSync(filePath, content, 'utf8');
}

function main() {
  const rawData = loadRawData();
  const dashboardData = buildDashboardData(rawData);
  writeDashboardData(dashboardData, outputPath);
  const sourceLabel = fs.existsSync(sourcesDir)
    ? path.relative(process.cwd(), sourcesDir)
    : path.relative(process.cwd(), sourcePath);
  console.log(`Generated ${path.relative(process.cwd(), outputPath)} from ${sourceLabel}`);
}

main();
