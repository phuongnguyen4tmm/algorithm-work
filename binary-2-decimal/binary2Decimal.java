// dùng số mũ
public static int binary2Decimal(String input) {
    double result = 0;

    int strLen = input.length();
    for (int i = 0; i < strLen; i++) {
        double power = strLen - i - 1;
        double number = input.charAt(i) - 48;
        result += Math.pow(2, power) * number;
    }

    return (int) result;
}

// Dùng nhân đôi
public static int binary2Decimal2(String input) {
    int result = 0;

    for (char character :
            input.toCharArray()) {
        int number = character - 48;
        result = result * 2 + number;
    }

    return result;
}