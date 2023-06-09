#include <cmath> 
#include<vector> 

using std::vector; 
class Solution {
public: 
    bool isPossible(int x, int y, char number,vector<vector<char>>& board){ 

        for(int i=0;i<9;i++){ 
            if(board[x][i] == number){ 
                return false;
            }  

            if(board[i][y] == number){ 
                return false;
            }
        } 

        int xo = floor(x / 3) * 3; 
        int yo = floor(y / 3) * 3; 

        for(int i=0;i<3;i++){ 
            for(int j=0;j<3;j++){ 
                if (board[xo+j][yo+i] == number){ 
                    return false;
                }
            }
        }
        return true;
    } 

    bool solve(vector<vector<char>>& board){  

        for(int i=0;i<9;i++){ 
            for(int j=0;j<9;j++){ 
                if(board[i][j] == '.'){
                    for(char num='1';num<='9';num++){ 
                        if(isPossible(i,j,num,board)){ 
                            board[i][j] = num;
                            if(solve(board)){ 
                                return true;
                            } else {  
                                board[i][j] = '.';
                            }
                        }
                    }
                    return false;
                }  
            } 
        } 
        return true; 
    }
    void solveSudoku(vector<vector<char>>& board) { 
       solve(board);
    }
};