class Ball {
  x: number;
  y: number;
  health: number;
  maxHealth: number;

  constructor() {
    this.x = 0;
    this.y = 0;
    this.health = 50;
    this.maxHealth = 50;
  }

  moveBall(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
}

export { Ball };
