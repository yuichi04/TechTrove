import SaySomething from "./saySomthing";

const root: HTMLElement | null = document.getElementById("root");

// インスタンス化
const saySomthing = new SaySomething("Hello TypeScript");
saySomthing.sayText(root);
