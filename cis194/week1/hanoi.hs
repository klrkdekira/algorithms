type Peg = String
type Move = (Peg, Peg)
hanoi :: Integer -> Peg -> Peg -> Peg -> [Move]
hanoi 1 a b _ = [(a, b)]
hanoi n a b c = 
    stepOne ++ stepTwo ++ stepThree
    where
        stepOne = hanoi (n-1) a c b
        stepTwo = hanoi 1 a b c
        stepThree = hanoi (n-1) c b a
