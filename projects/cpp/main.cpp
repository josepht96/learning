#include <iostream>
#include <chrono>
#include <ctime>
using namespace std::chrono;

int main()
{
    auto start = high_resolution_clock::now();
    for (int i = 0; i < 1000000000; i++) {
        int v = i * 100;
        int v2 = v / 99;
        //std::cout << v2 << "\n";
    }
    auto stop = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(stop - start);
 
    // To get the value of duration use the count()
    // member function on the duration object
    std::cout << duration.count() << "\n";
    return 0;
}