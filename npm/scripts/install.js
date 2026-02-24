const https = require("https");
const fs = require("fs");
const path = require("path");
const os = require("os");

const pkg = require("../package.json");
const version = pkg.version;

const platformMap = { linux: "linux", darwin: "darwin", win32: "windows" };
const archMap = { x64: "amd64", arm64: "arm64" };

const goos = platformMap[os.platform()];
const goarch = archMap[os.arch()];

if (!goos || !goarch) {
  console.error(`Unsupported platform: ${os.platform()}/${os.arch()}`);
  process.exit(1);
}

const ext = os.platform() === "win32" ? ".exe" : "";
const binaryName = `openclaude-${goos}-${goarch}${ext}`;
const url = `https://github.com/johmara/openclaude/releases/download/v${version}/${binaryName}`;
const dest = path.join(__dirname, "..", "bin", binaryName);

function download(url, dest, redirects) {
  if (redirects === undefined) redirects = 5;
  if (redirects === 0) {
    console.error("Too many redirects");
    process.exit(1);
  }

  https
    .get(url, { headers: { "User-Agent": "openclaude-npm" } }, (res) => {
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        // Follow redirect (GitHub redirects to S3)
        const redirectUrl = res.headers.location;
        // Use https or http depending on redirect URL
        const mod = redirectUrl.startsWith("https") ? https : require("http");
        mod
          .get(redirectUrl, { headers: { "User-Agent": "openclaude-npm" } }, (res2) => {
            if (res2.statusCode !== 200) {
              console.error(`Download failed: HTTP ${res2.statusCode}`);
              process.exit(1);
            }
            const file = fs.createWriteStream(dest);
            res2.pipe(file);
            file.on("finish", () => {
              file.close();
              fs.chmodSync(dest, 0o755);
              console.log(`Downloaded openclaude v${version} (${goos}/${goarch})`);
            });
          })
          .on("error", (err) => {
            console.error("Download error:", err.message);
            process.exit(1);
          });
        return;
      }

      if (res.statusCode !== 200) {
        console.error(`Download failed: HTTP ${res.statusCode}`);
        process.exit(1);
      }

      const file = fs.createWriteStream(dest);
      res.pipe(file);
      file.on("finish", () => {
        file.close();
        fs.chmodSync(dest, 0o755);
        console.log(`Downloaded openclaude v${version} (${goos}/${goarch})`);
      });
    })
    .on("error", (err) => {
      console.error("Download error:", err.message);
      process.exit(1);
    });
}

// Ensure bin directory exists
const binDir = path.dirname(dest);
if (!fs.existsSync(binDir)) {
  fs.mkdirSync(binDir, { recursive: true });
}

download(url, dest);
