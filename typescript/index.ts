import tasks from './tasks';

function main(): void {
  console.log(tasks.evenOrOdd(15));
  console.log(tasks.coutingSheep([true, true, false, false, true]));
  console.log(tasks.countTheMonkeys(13));
  console.log(tasks.schoolPaperwork(3, 5));
  console.log(tasks.defeatingDragons(11, 6));
  console.log(tasks.polishAlphabet('Jędrzej Błądziński'));
  console.log(tasks.findAll([1, 2, 2, 3, 3, 2, 4, 5, 2], 2));
  console.log(
    tasks.sumOfMinimums([
      [1, 2, 3, 4, 5],
      [5, 6, 7, 8, 9],
      [20, 21, 34, 56, 100]
    ])
  );
  //tasks.greenBottles(8)
}

main();
