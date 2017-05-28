toDigits :: Integer -> [Integer]
toDigits x
    | x > 0 = toDigits (div x 10) ++ [mod x 10]
    | otherwise = []

toDigitsRev :: Integer -> [Integer]
toDigitsRev = reverse . toDigits

doubleEven :: Integer -> Integer -> Integer
doubleEven i x
    | even i = x * 2
    | otherwise = x

doubleEveryOther :: [Integer] -> [Integer]
doubleEveryOther = zipWith doubleEven [1..]

sumDigits :: [Integer] -> Integer
sumDigits = sum . map (sum . toDigits)

validate :: Integer -> Bool
validate num = 
    sum' num `mod` 10 == 0
    where
        sum' = sumDigits . doubleEveryOther . toDigitsRev

