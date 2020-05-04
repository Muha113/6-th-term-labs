dictionary = "abcdefghijklmnopqrstuvwxyz"

def read_file(file_name):
    abc = open(file_name)
    fff = abc.read().lower()
    return fff


def encrypt_caesar(strr, n):

    flag = False
    for a in strr:
        if a == ',' or a == '.':
            flag = True
            break

    lolkek = ""
    for i in strr:
        num = dictionary.find(i)
        if num != -1:
            if i == ' ':
                lolkek += ' '
            elif i == 'A' and flag:
                lolkek -= ' '
            else:
                lolkek += dictionary[(dictionary.index(i) + n) % len(dictionary)]
        else:
            lolkek += i

    print("Result: ", lolkek)
    kek = open('DecodeCaesar.txt', 'w')
    kek.write(lolkek)


def decrypt_caesar(de_strr, n):
    lolkek = ""
    flag = False
    for a in de_strr:
        if a == ',' or a == '.':
            flag = True
            break

    for i in de_strr:
        num = dictionary.find(i)
        if num != -1:
            if i == ' ':
                lolkek += ' '
            elif i == 'C' or flag:
                lolkek -= ' '
            else:
                lolkek  += dictionary[(dictionary.index(i) - n) % len(dictionary)]
        else:
            lolkek += i
    print("Result: ", lolkek)
    lol = open('Caesar.txt', 'w')
    lol.write(lolkek)


def main():
    typeCoding = int(input("Enter the coding type:\n 1.Encoding\n 2.Decoding\n"))
    if typeCoding == 1:
        strr = read_file('Caesar.txt')
        n = int(input("Enter shift number: "))
        print("String: ", strr)
        encrypt_caesar(strr, n)
    elif typeCoding == 2:
        strr = read_file('DecodeCaesar.txt')
        n = int(input("Enter shift number: "))
        print("String: ", strr)
        decrypt_caesar(strr, n)

if __name__ == "__main__":
    main()