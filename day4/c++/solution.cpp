#include <iostream>
#include <fstream>
#include <map>
#include <regex>
#include <vector>

using namespace std;

const string FILENAME = "../input";
const regex SPLIT("\\n\\n");

string readFile(string filename) {
    string line;
    string data = "";
    ifstream inputFile(filename);
    while(getline(inputFile, line)) {
        data += line + "\n";
    }
    return data;
}

vector<string> splitPassports(string data) {
    vector<string> passports;
    sregex_token_iterator i(data.begin(), data.end(), SPLIT, -1);
    sregex_token_iterator end;
    while(i != end) {
        passports.push_back(*i);
        *i++;
    }
    return passports;
}

int countValidPassports(vector<string> passports) {
    int valid = 0;
    smatch match;
    for(auto i = passports.begin(); i != passports.end(); i++) {
        string passport = *i;
        if(passport.find("byr") != string::npos &&passport.find("iyr") != string::npos &&passport.find("eyr") != string::npos &&passport.find("hgt") != string::npos &&passport.find("hcl") != string::npos &&passport.find("ecl") != string::npos &&passport.find("pid") != string::npos) valid++;
    }
    return valid;
}

int main() {
    string data = readFile(FILENAME);
    vector<string> passports = splitPassports(data);
    cout << countValidPassports(passports) << endl;
}