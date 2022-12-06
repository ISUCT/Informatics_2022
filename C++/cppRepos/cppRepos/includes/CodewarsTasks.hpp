#ifndef _CodewarsTasks_HPP_
    
#define _CodewarsTasks_HPP_
#include <string>
#include <vector>
#include <list>
#include <map>
    
using namespace std;

auto EvenOrOddTask(unsigned int num) -> string;

auto CountingSheepTask(vector<bool>& sheeps) -> int;
    
auto CountingSheepTask(list<bool>& sheeps) -> int;

auto CountingSheepTask(bool* sheeps, unsigned int size) -> int;

auto CountTheMonkeysTask(int n) -> int*;

auto SchoolPaperworkTask(int n, int m) -> int;

auto IsHeGonnaSurvie(int countBullets, int countDragons) -> bool;

auto PolishAlphabetTask(wstring str) -> string;

auto FindAllOccurrencesOfAnElementInAnArrayTask(vector<int>& arr, int num) -> vector<int>&;

auto FindAllOccurrencesOfAnElementInAnArrayTask(int* arr, unsigned int size, int num) -> int*;

auto SumOfMinimums(vector<vector<int>> arr2d) -> int;

auto SumOfMinimums(int** arr2d, unsigned int m, unsigned int n) -> int;

#endif

