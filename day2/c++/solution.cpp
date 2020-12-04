#include <iostream>
#include <vector>
#include <fstream>
#include <regex>
#include <string>
#include <algorithm>

using namespace std;

const string FILENAME = "../input";
const int SLED = 1;
const int TOB = 2;

const regex R("(\\d{1,2})-(\\d{1,2})\\s(\\w):\\s(\\b\\w+\\b)");

int countLetters(string inputString, char inputChar) {
	int count = 0;

	for (int i : inputString) {
		if (i == inputChar) {
			count++;
		}
	}
	return count;
}

int validSled(string passwd, char letter, int min, int max) {
        int appears = countLetters(passwd, letter);

        if (appears >= min && appears <= max) {
                return 1;
        }
	return 0;
}

int validTob(string passwd, char letter, int first, int second) {
	// Make sure everything is in range
	int length = passwd.length();
	if (length < first || length < second) {
		first = second = 1;
	}

	// Convert from one index to zero index
	first -= 1;
	second -= 1;

	if ((passwd[first] == letter) != (passwd[second] == letter)) {
		return 1;
	}
	return 0;
}

int validate(string filename, int corp) {
	int count = 0;
	string line;

	smatch match;

	ifstream inputFile(filename);
	while(getline(inputFile, line)) {
		// Check for empty lines
		if (line == "") {
			continue;
		}

		regex_search(line, match, R);
		int min = stoi(match.str(1));
		int max = stoi(match.str(2));
		char letter = match.str(3)[0];

		switch(corp) {
			case SLED:
				count += validSled(match.str(4), letter, min, max);
				break;
			case TOB:
				count += validTob(match.str(4), letter, min, max);
				break;
		}
	}
	return count;
}

int main() {
	int validS = validate(FILENAME, SLED);
	int validT = validate(FILENAME, TOB);
	cout << validS << endl << validT << endl;
}
