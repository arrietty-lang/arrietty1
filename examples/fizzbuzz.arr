int main() {
    for (i:=1; i<101; i=i+1) {
        if (i%15 == 0) {
            print( itos(i) + " fizzbuzz" + "\n" );
        } else if (i%3 == 0) {
            print( itos(i) + " fizz" + "\n" );
        } else if (i%5 == 0) {
            print( itos(i) + " buzz" + "\n" );
        } else {
            print( itos(i) + "\n" );
        }
    }
    return 0;
}
