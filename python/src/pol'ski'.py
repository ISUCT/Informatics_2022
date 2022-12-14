# https://www.codewars.com/kata/57ab2d6072292dbf7c000039
abs = {"ą": "a",
       "ć": "c",
       "ę": "e",
       "ł": "l",
       "ń": "n",
       "ó": "o",
       "ś": "s",
       "ź": "z",
       "ż": "z"}

word = "Jędrzej Błądziński"

for i in abs.keys():
    word = word.replace(i, abs[i])
print(word)

