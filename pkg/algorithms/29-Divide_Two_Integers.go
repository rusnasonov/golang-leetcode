package algorithms


func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func div(dividend, divisor int) int {
    isNegative := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)
    quotient := 0
    dividend = abs(dividend)
    divisor = abs(divisor)
    maxInt := (1 << 31) - 1
    minInt := (-1 << 31)
    for dividend >= divisor {
        shift := 0
        for dividend >= (divisor << shift) {
            shift += 1
        }
        quotient += (1 << (shift - 1))
        dividend -= (divisor << (shift - 1))
    }
    if isNegative {
        quotient  = quotient * -1
    }
    if quotient > maxInt {
        quotient = maxInt
    }
    if quotient < minInt {
        quotient = minInt
    }
    return quotient
}

func divide(dividend int, divisor int) int {
    return div(dividend, divisor)
}
