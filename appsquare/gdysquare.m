function rmat = gdysquare(mat)
%此函数用于矩阵求平方和
%输入1个矩阵，输出矩阵的平方和
[r,c]=size(mat);
sum = 0;
for i = 1:r
   for j = 1:c
    sum = sum + mat(i,j)*mat(i,j);
   end
end
rmat = sum;
