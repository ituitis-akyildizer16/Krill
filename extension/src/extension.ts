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
  });
}

function workspaceRoot(): string {
  return vscode.workspace.workspaceFolders?.[0]?.uri.fsPath ?? ".";
}

export function activate(context: vscode.ExtensionContext): void {
  context.subscriptions.push(
    vscode.commands.registerCommand("krill.ask", async () => {
      const editor = vscode.window.activeTextEditor;
      if (!editor) {
        return vscode.window.showWarningMessage("Open a file first");
      }
      const file = editor.document.uri.fsPath;
      const question = await vscode.window.showInputBox({
        prompt: "Question about this file",
        placeHolder: "why is this function slow?",
      });
      if (!question) {
        return;
      }
      const answer = await vscode.window.withProgress(
        { location: vscode.ProgressLocation.Notification, title: "krill" },
        () => runKrill(["ask", question, "-f", file]),
      );
      const panel = vscode.window.createWebviewPanel(
        "krill.answer", "krill", vscode.ViewColumn.Beside,
        { enableScripts: false },
      );
      panel.webview.html = renderMarkdown(answer);
    }),

    vscode.commands.registerCommand("krill.review", async () => {
      const summary = await vscode.window.withProgress(
        { location: vscode.ProgressLocation.Notification, title: "krill" },
        () => runKrill(["review", "--staged"]),
      );
      vscode.window.showInformationMessage(summary);
    }),

    vscode.commands.registerCommand("krill.suggest", async () => {
      const intent = await vscode.window.showInputBox({
        prompt: "What do you want to do?",
      });
      if (!intent) {
        return;
      }
      const cmd = await runKrill(["suggest", "--inline", intent]);
      vscode.env.clipboard.writeText(cmd);
      vscode.window.showInformationMessage(`Copied: ${cmd}`);
    }),
  );
}

function renderMarkdown(text: string): string {
  const body = text
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/\n/g, "<br/>");
  return `<!DOCTYPE html>
<html><body style="font-family: sans-serif; padding: 1rem">
<pre>${body}</pre>
</body></html>`;
}

export function deactivate(): void {}