int main() {
    for (i:=0; i<100; i=i+1) {
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