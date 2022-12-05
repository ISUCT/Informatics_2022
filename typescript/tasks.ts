import { IntRange } from './types';
import { numsToString, mapping } from './variables';

const tasks = {
  evenOrOdd: (number: number): 'Even' | 'Odd' => {
    return number % 2 == 0 ? 'Even' : 'Odd';
  },
  coutingSheep: (sheeps: boolean[]): number => {
    return sheeps.reduce((acc, sheep) => {
      return sheep ? (acc += 1) : acc;
    }, 0);
  },
  countTheMonkeys: (number: number): number[] => {
    if (number <= 0) {
      return [];
    }

    const result: number[] = [];
    for (let i = 1; i <= number; i++) {
      result.push(i);
    }
    return result;
  },
  greenBottles: (number: IntRange<1, 10>): void => {
    for (let i = number; i > 0; i--) {
      const bottleNumbers =
        numsToString[i - 1][0].toUpperCase() + numsToString[i - 1].slice(1);
      if (i >= 1) {
        console.log(`${bottleNumbers} green bottles hanging on the wall,`);
        console.log(`${bottleNumbers} green bottles hanging on the wall,`);
      }

      if (i == 1) {
        console.log('If that one green bottle should accidentally fall,');
        console.log("There'll be no green bottles hanging on the wall.");
      } else {
        console.log('And if one green bottle should accidentally fall,');
        console.log(
          `There'll be ${
            numsToString[i - 2]
          } green bottles hanging on the wall. \n`
        );
      }
    }
  },
  schoolPaperwork: (n: number, m: number): number => {
    if (n < 0 || m < 0) {
      return 0;
    }
    return n * m;
  },
  defeatingDragons: (bulletsCount: number, dragonsCount: number): boolean => {
    return Math.floor(bulletsCount / 2) >= dragonsCount ? true : false;
  },
  polishAlphabet: (str: string): string => {
    return str.replace(/[^a-zA-Z]/g, (c) => mapping[c] ?? c);
  },
  findAll: (arr: number[], num: number): number[] => {
    return arr.reduce<number[]>(
      (acc, currentValue, index) =>
        num === currentValue ? [...acc, index] : acc,
      []
    );
  },
  sumOfMinimums: (arr: number[][]): number => {
    return arr.reduce((acc, currentArr) => acc + Math.min(...currentArr), 0);
  }
};

export default tasks;
