#include <iostream>
#include <fstream>
#include <string>

using namespace std;

const string FILENAME = "../input";

int countTrees(int right, int down) {
    int trees = 0;
    int hCount = 0;
    int vCount = 0;
    int location = 0;
    string line;
    ifstream inputFile(FILENAME);
    while(getline(inputFile, line)) {
        // Go down multiples lines at once
        if (vCount % down != 0) {
            vCount++;
            continue;
        }

        // Check for trees
        if (line[hCount % line.length()] == '#') {
            trees++;
        }

        hCount += right;
        vCount ++;
    }
    return trees;
}

int main() {
    long nums[] = {
        countTrees(1,1),
        countTrees(3,1),
        countTrees(5,1),
        countTrees(7,1),
        countTrees(1,2)
    };
    cout << nums[0] << endl << nums[1] << endl << nums[2] << endl << nums[3] << endl <<
        nums[4] << endl;
    long product = nums[0] * nums[1] * nums[2] * nums[3] * nums[4];
    cout << product << endl;
}