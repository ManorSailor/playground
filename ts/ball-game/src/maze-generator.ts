type Wall = {
  top: boolean;
  right: boolean;
  bottom: boolean;
  left: boolean;
};

type Cell = {
  x: number;
  y: number;
  walls: Wall;
};

type Maze = {
  rows: number;
  cols: number;
  cells: Cell[][];
};

interface IMazeGenerator {
  readonly rows: number;
  readonly cols: number;

  generateMaze: () => Maze;
}

class MazeGenerator implements IMazeGenerator {
  readonly rows;
  readonly cols;
  private cells: Cell[][];
  private visitedCells: Set<Cell>;

  constructor(rows: number = 10, cols: number = 10) {
    this.rows = rows;
    this.cols = cols;
    this.cells = this.generateCells();
    this.visitedCells = new Set();
  }

  generateMaze() {
    const generate = (startingCell: Cell) => {
      this.visitedCells.add(startingCell);

      let unvisitedCells = this.getUnvisitedCells(startingCell);

      while (unvisitedCells.length) {
        const idx = Math.floor(Math.random() * unvisitedCells.length);
        const nextCell = unvisitedCells[idx];

        this.removeWall(startingCell, nextCell);
        generate(nextCell);

        unvisitedCells = this.getUnvisitedCells(startingCell);
      }
    };

    generate(this.cells[0][0]);

    return {
      rows: this.rows,
      cols: this.cols,
      cells: this.cells,
    };
  }

  private generateCells() {
    const cells: Cell[][] = [];

    for (let r = 0; r < this.rows; r++) {
      const row: Cell[] = [];
      for (let c = 0; c < this.cols; c++) {
        row.push({
          x: c,
          y: r,
          walls: {
            top: true,
            bottom: true,
            right: true,
            left: true,
          },
        });
      }
      cells.push(row);
    }

    return cells;
  }

  private getUnvisitedCells(cell: Cell) {
    const { x, y } = cell;
    const cells = this.cells;

    const offsets = [
      [-1, 0],
      [1, 0],
      [0, -1],
      [0, 1],
    ];

    return offsets.reduce((unvisitedCells, [offsetX, offsetY]) => {
      const px = offsetX + x;
      const py = offsetY + y;

      if (px >= 0 && px < this.cols && py >= 0 && py < this.rows) {
        const cell = cells[py][px];

        if (!this.visitedCells.has(cell)) {
          unvisitedCells.push(cell);
        }
      }

      return unvisitedCells;
    }, [] as Cell[]);
  }

  private removeWall(cellA: Cell, cellB: Cell) {
    const dx = cellA.x - cellB.x;
    const dy = cellA.y - cellB.y;

    if (dx === 1) {
      cellA.walls.left = false;
      cellB.walls.right = false;
    } else if (dx === -1) {
      cellA.walls.right = false;
      cellB.walls.left = false;
    }

    if (dy === 1) {
      cellA.walls.top = false;
      cellB.walls.bottom = false;
    } else if (dy === -1) {
      cellA.walls.bottom = false;
      cellB.walls.top = false;
    }
  }
}

export { MazeGenerator, type Cell, type Maze };
