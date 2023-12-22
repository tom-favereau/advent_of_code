# AOS 2023

     __,_,_,___)          _______
    (--| | |             (--/    ),_)       ,_)
       | | |  _ ,_,_        |     |_ ,_' , _|_,_,_, _  ,
     __| | | (/_| | (_|     |     | || |/_)_| | | |(_|/_)___,
    (      |___,   ,__|     \____)  |__,          |__,

                            |                         _...._
                         \  _  /                    .::o:::::.
                          (\o/)                    .:::'''':o:.
                      ---  / \  ---                :o:_    _:::
                           >*<                     `:}_>()<_{:'
                          >0<@<                 @    `'//\\'`    @
                         >>>@<<*              @ #     //  \\     # @
                        >@>*<0<<<           __#_#____/'____'\____#_#__
                       >*>>@<<<@<<         [__________________________]
                      >@>>0<<<*<<@<         |=_- .-/\ /\ /\ /\--. =_-|
                     >*>>0<<@<<<@<<<        |-_= | \ \\ \\ \\ \ |-_=-|
                    >@>>*<<@<>*<<0<*<       |_=-=| / // // // / |_=-_|
      \*/          >0>>*<<@<>0><<*<@<<      |=_- |`-'`-'`-'`-'  |=_=-|



benchmark done with mac intel

cpu: Intel(R) Core(TM) i5-1030NG7 CPU @ 1.10GHz

## day1
done part one and part two, easy
I learn to pars an input

benchmark :

| Part1    | Part2    | 
|----------|----------|
| 1.296 ms | 1.430 ms |

## day2
done part one and part two, easy

| Part1    | Part2    | 
|----------|----------|
| 0.129 ms | 0.158 ms |

## day3
done part one and part two, medium

| Part1    | Part2    | 
|----------|----------|
| 1,105 ms | 0.792 ms |


## day4 
done part one and part tow, easy
I must read more attentively the subject 

| Part1    | Part2    | 
|----------|----------|
| 0.488 ms | 0.451 ms |


## day5 
done part1 and part2, difficult

nice solution

| Part1    | Part2    | 
|----------|----------|
| 0.060 ms | 0.082 ms |

## day6 
done part1 and part2, very easy

| Part1    | Part2    | 
|----------|----------|
| 0.012 ms | 0.012 ms |

## day7 
done part1 and part2
I learn how to use the sort package 

| Part1    | Part2    | 
|----------|----------|
| 2.802 ms | 2.773 ms |

## day8

done part1 and part2
it was not supposed to work

| Part1    | Part2    |
|----------|----------|
| 0.653 ms | 5.624 ms |

## day9

done part1 and part2

I learn a new way to extend additive sequence


| Part1    | Part2    | 
|----------|----------|
| 0.297 ms | 0.309 ms |

## day10

difficile j'ai perdu trop de temps dans un code ilisible que j'arrvé plus à refaire
pensé à faire de commit régulier pour pas perdre son code

| Part1    | Part2    | 
|----------|----------|
| 0.604 ms | 0.609 ms |

## day 11 

Done part1 and part2
not realy optimised 

| Part1    | Part2    | 
|----------|----------|
| 101.4 ms | 105.2 ms |


## day 12

Done part1 and part2
part 1 avec élagation 
part 2 avec programation dynamique, je sais pas si tout les paramètre mémorisé sont utile

| Part1   | Part2   | 
|---------|---------|
| 15.34ms | 32.9 ms |

## day 13

Done part1 and part2
le cours de management permet de trouver de très bonne astuce comme 
s'aréter s'il n'y a qu'une diférence pour la partie 2

| Part1   | Part2   | 
|---------|---------|
| 4.03 ms | 4.56 ms |

## day 14

Done part1 and part2
HORRIBLE 
j'ai pas aimé du tout, ne pas tenir compte de ce problème dans la notation, je suis même plus sur qu'il soit fonctionel
et je ne veux plus jamais y toucher.



## day 15

Done part1 and part2
Implémentation d'une hashmap, on doit sans doute pouvoir utilisé une hashmap buildin


| Part1   | Part2   | 
|---------|---------|
| 0.92 ms | 2.84 ms |

## day 16

Done part1 and part2
On récupère le jour 10 part 1 on fait un DFS avec des direction

| Part1   | Part2  | 
|---------|--------|
| 3.98 ms | 716 ms |


## day 17

Done part1 and part2
on fait un dijikstra en 4D en définissant les états comme i, j, position, longueur parcurus dans cette direction
et on vérifie les condition au borne c'est a dire, 0 < i < X, 0 < j < Y et 0 < longueur < 4 etc ..
idem pour la partie 4

| Part1   | Part2   | 
|---------|---------|
| 4760 ms | 4590 ms |

## day 18

Done part1 and part2
on reprend la partie 2 du jour 10 formule de l'arpenteur + formule de pick

| Part1    | Part2    | 
|----------|----------|
| 0.132 ms | 0.273 ms |


## day 19

Done part1 and part2
très enuyeux, juste on problème avec du parsing, rien d'intéréssant. on reprend les intervalle du jour 5 rien de plus a 
noter 


| Part1   | Part2   | 
|---------|---------|
| 8.50 ms | 2.52 ms |

## day 20

Done part1 and part2
jour difficile en grande partie car il fallait regarder l'input, j'ai discuté avec victor ce qui m'a permis d'avoir la bonne aproche 
comme j'avais un bug que je ne trouvais pas (un != à la place d'un ==) je suis allé sur disocrd et j'ai vue les graphes 
j'ai donc résolue le problème comme ça. une fois le problème validé j'ai repris mon programme et coriggé l'erreur rapidement.
on trouve les longueur des cycle et on fait le ppcm (un produit suffit car il sont premier)



| Part1   | Part2   | 
|---------|---------|
| 10.4 ms | 34.4 ms |

## day 21

Done part1 and part2
très dur, j'ateint mes limites, je n'ai pas trouvé d'algorithme, 
au départ je n'ai rien trouvé après être allé sur reddit j'ai vue une animation 
ma première approche a était de considéré le diamand calculer sont aire
donné par la formule 2*(x**2)-2*x (somme des nombres impaire * 2 - la ligne du millieux)
après avoir galéré pendant 5 bonnes heure j'ai finis par validé. 
puisque c'est polynomial on pouvais aussi interpoler sur les 3 première valeur (encore fallait il le voir venir) 
par conséquent après avoir visité reddit j'ai implémenter les polynome de lagrange pour avoir un code
néamoins je ne considère pas ça comme un vrais algorithme et je serais curieux de savoir ce qu'on fait les autre.
**pour l'intant je ne vais pas voir je cherche encore**




## day 22

Done part1 and part2
assez facile, rien de bien méchant, l'optimiser est peux être plus dur, mais à ce stade du mois je ne cherche plus trop a optimiser


| Part1  | Part2  | 
|--------|--------|
| 152 ms | 686 ms |