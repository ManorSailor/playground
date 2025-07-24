import { Ball } from "./ball";
import { MazeGenerator, type Maze } from "./maze-generator";
import "./style.css";

const ball = document.querySelector("span.ball")! as HTMLSpanElement;
const playground = document.querySelector("#app")! as HTMLDivElement;
const hole = document.querySelector(".hole")! as HTMLDivElement;

const mazeGen = new MazeGenerator(30, 30)
// const ball = new Ball()

playground.addEventListener("mousemove", gameLoop);

function gameLoop(e: MouseEvent) {
  const holeRect = hole.getBoundingClientRect();
  const ballRect = ball.getBoundingClientRect();
 
  if (isColliding(holeRect, ballRect)) {
    resetBall();
    playground.removeEventListener("mousemove", gameLoop);
    setTimeout(() => playground.addEventListener("mousemove", gameLoop), 3000);
  } else {
    moveBall(e.clientX, e.clientY);
  }
}

function moveBall(posX: number, posY: number) {
  ball.style.transform = `translate3d(${posX}px, ${posY}px, 0)`;
}

function resetBall() {
  ball.style.transform = "translate3d(0, 0, 0)";
}

function isColliding(a: DOMRect, b: DOMRect) {
  return !(
    a.y + a.height < b.y ||
    a.y > b.y + b.height ||
    a.x + a.width < b.x ||
    a.x > b.x + b.width
  );
}

function renderMaze({ cols, cells: grid }: Maze) {
  const mazeContainer = document.getElementById("maze")!;
  mazeContainer.innerHTML = "";

  mazeContainer.style.gridTemplateColumns = `repeat(${cols}, 1fr)`;

  grid.forEach((row) => {
    row.forEach((cell) => {
      const cellDiv = document.createElement("div");

      const { top, right, bottom, left } = cell.walls;

      cellDiv.style.borderTop = top ? "2px solid black" : "none";
      cellDiv.style.borderRight = right ? "2px solid black" : "none";
      cellDiv.style.borderBottom = bottom ? "2px solid black" : "none";
      cellDiv.style.borderLeft = left ? "2px solid black" : "none";

      mazeContainer.appendChild(cellDiv);
    });
  });
}

renderMaze(mazeGen.generateMaze());
