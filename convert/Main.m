function [Main]=Main()

%计算类型参数
arg_list = argv();
type= str2num(arg_list{1});
fprintf('Type = %d\n',type)

%静态节点冲剪参数
%%%%%%%%%%%EffectiveT(有效厚度限制)
% EffectiveT=1.75;
dir='EffectiveT.csv';
EffectiveT=csvread(dir);


%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%核心计算内容%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%
E='';
JCNout=JCN(1,EffectiveT);
if strcmp(JCNout,'Done')
    fprintf('Jointcan File for Static Analysis Generated%s\n',E)
end

Main='Done';