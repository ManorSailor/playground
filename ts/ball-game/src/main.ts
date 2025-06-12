import "./style.css";

const ball = document.querySelector("span.ball")! as HTMLSpanElement;
const playground = document.querySelector("#app")! as HTMLDivElement;
const hole = document.querySelector(".hole")! as HTMLDivElement;

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
