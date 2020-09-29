import * as vscode from "vscode";
import { execFile } from "child_process";
import { promisify } from "util";

const exec = promisify(execFile);

function binary(): string {
  return vscode.workspace
    .getConfiguration("krill")
    .get<string>("binaryPath", "krill");
}

function runKrill(args: string[], input?: string): Promise<string> {
  return new Promise((resolve, reject) => {
    const child = exec(binary(), args, { cwd: workspaceRoot() });
    child.then(({ stdout }) => resolve(stdout.trim())).catch(reject);
