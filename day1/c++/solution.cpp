#include <iostream>
#include <fstream>
#include <string>
using namespace std;

const string FILENAME = "../input";


int getFileSize(string filename) {
	string text;
	int count;

	ifstream inputFile(filename);
	while(getline(inputFile, text)) {
		count++;
	}
	return count;
}

int* getTwoNumbers(int* array, int length) {
	static int numbers[2];
	for(int i = 0; i < length; i++) {
		for(int j = i; j < length; j++) {
			if(array[i] + array[j] == 2020) {
				numbers[0] = array[i];
				numbers[1] = array[j];
				return numbers;
			}
		}
	}
	return numbers;
}

int* getThreeNumbers(int* array, int length) {
	static int numbers[3];
	for(int i = 0; i < length; i++) {
		for(int j = i; j < length; j++){
			for(int k = j; k < length; k++) {
				if(array[i] + array[j] + array[k] == 2020) {
					numbers[0] = array[i];
					numbers[1] = array[j];
					numbers[2] = array[k];
					return numbers;
				}
			}
		}
	}
	return numbers;
}

int main() {
	string text;
	int count = 0;

	int length = getFileSize(FILENAME);
	int* numArray = new int[length];

	ifstream inputFile(FILENAME);
	while(getline(inputFile, text)) {
		numArray[count] = stoi(text);
		count++;
	}

	int* twoSumNums = getTwoNumbers(numArray, length);
	int* threeSumNums = getThreeNumbers(numArray, length);
	cout << twoSumNums[0] * twoSumNums[1] << endl << threeSumNums[0] * threeSumNums[1] * threeSumNums[2] << endl;
	return 0;
}
