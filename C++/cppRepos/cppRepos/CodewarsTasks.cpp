#include "includes/CodewarsTasks.hpp"

string EvenOrOddTask(unsigned int num) {
    string answ;
    return num % 2 == 0 ? answ = "Even" : answ = "Odd";
}

int CountingSheepTask(vector<bool>& sheeps) {
    int count = 0;
    for (auto it : sheeps) 
        if (it) ++count;
    return count;
}

int CountingSheepTask(list<bool>& sheeps) {
    int count = 0;
    for (auto it : sheeps) 
        if (it) ++count;
    return count;
}

vector<int> CountTheMonkeysTask(int n) {
    vector<int> arr;
    arr.resize(n);
    for (int i = 0; i < n; ++i) 
        arr[i] = i + 1;
    return arr;
}

int SchoolPaperworkTask(int n, int m) {
    int res;
    return n < 0 || m < 0 ? res = 0 : res = n * m;
}

bool IsHeGonnaSurvie(int countBullets, int countDragons) {
    return countBullets >= countDragons * 2;
}

string PolishAlphabetTask(wstring str) {
    string result;
    for (auto ch : str) {
        switch (ch) {
        case L'ą': result += 'a'; break;
        case L'ć': result += 'c'; break;
        case L'ę': result += 'e'; break;
        case L'ł': result += 'l'; break;
        case L'ń': result += 'n'; break;
        case L'ó': result += 'o'; break;
        case L'ś': result += 's'; break;
        case L'ź': 
        case L'ż': result += 'z'; break;
        default: result += (char)ch;
        }
    }
    return result;
}

vector<int> FindAllOccurrencesOfAnElementInAnArrayTask(vector<int>& arr, int num) {
    vector<int> res;
    for (int i = 0; i < arr.size(); ++i)
        if (arr[i] == num)
            res.push_back(i);
    return res;
}

int SumOfMinimums(vector<vector<int>> arr2d) {
    int res = 0;
    for (auto it : arr2d) {
        res += *min(it.cbegin(), it.cend());
    }
    return res;
}