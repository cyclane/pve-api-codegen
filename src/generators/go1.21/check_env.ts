export default async function () {
  // check go executable and version
  const command = new Deno.Command("go", {
    args: ["version"],
    stdout: "piped",
    stderr: "piped",
  });
  const child = command.spawn();
  const status = await child.status;
  const output = await child.output();

  const td = new TextDecoder();
  const stdout = td.decode(output.stdout).trim();
  const stderr = td.decode(output.stderr).trim();
  if (!status.success) {
    throw new Error(
      `go version failed:\n\tSTDOUT: ${stdout}\n\tSTDERR: ${stderr}`,
    );
  }
  if (!stdout.match(/^go version go1\.21\.\d+.*$/)) {
    throw new Error(
      `go version failed: expected 'go version go1.21...', got: '${stdout}'`,
    );
  }
}
