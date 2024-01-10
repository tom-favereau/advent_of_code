
module Main where 

import Data.List.Split (splitOn)
import System.Environment
import Data.Char (isDigit, ord)


part1 :: [String] -> Int 
part1 ll = sum (map aux ll) where 
    aux s = 10*first s + first (reverse s) where 
        first (h:t) = if isDigit h then ord h - 48 else first t 
        first _ = 0

part2 :: [String] -> Int 
part2 ll = sum (map aux ll) where
    aux s = 10*first s + last (reverse s) where 
        first ('o':'n':'e':_) = 1 
        first ('t':'w':'o':_) = 2 
        first ('t':'h':'r':'e':'e':_) = 3 
        first ('f':'o':'u':'r':_) = 4 
        first ('f':'i':'v':'e':_) = 5 
        first ('s':'i':'x':_) = 6 
        first ('s':'e':'v':'e':'n':_) = 7 
        first ('e':'i':'g':'h':'t':_) = 8 
        first ('n':'i':'n':'e':_) = 9 
        first (h:t) = if isDigit h then ord h - 48 else first t 
        first _ = 0
        last ('e':'n':'o':_) = 1 
        last ('o':'w':'t':_) = 2 
        last ('e':'e':'r':'h':'t':_) = 3 
        last ('r':'u':'o':'f':_) = 4 
        last ('e':'v':'i':'f':_) = 5 
        last ('x':'i':'s':_) = 6 
        last ('n':'e':'v':'e':'s':_) = 7 
        last ('t':'h':'g':'i':'e':_) = 8 
        last ('e':'n':'i':'n':_) = 9 
        last (h:t) = if isDigit h then ord h - 48 else first t 
        last _ = 0

main :: IO ()
main =  do 
    args <- getArgs
    case args of 
        [file_path] -> do 
            let content = readFile file_path 
            content >>= print . part1 . lines 
            content >>= print . part2 . lines 



