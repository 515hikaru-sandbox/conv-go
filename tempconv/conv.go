package tempconv

// CToF は摂氏温度を華氏温度に変換する
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏温度を摂氏温度に変換する
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
