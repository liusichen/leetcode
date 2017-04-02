#include<stdio.h>
#include<stdlib.h>
#include<string.h>

char *convert(char *s, int rownum);
int main(){
    printf("please enter the length of your string:");
    int leng;
    scanf("%d",&leng);
    char *str = (char *)malloc(sizeof(char)*(leng+1));
    char *res = (char *)malloc(sizeof(char)*(leng+1));
    scanf("%s",str);	
    int rownum,i;
    scanf("%d",&rownum);
    res = convert(str,rownum);
    for(i = 0; res[i]!= '\0'; i++){
	printf("%c",res[i]);
    }
    printf("\n");
    return 0;
}

char *convert(char *s, int rownum){
    int slen = strlen(s);
    char *res = (char *)malloc(sizeof(char)*(slen+1));
    bzero(res,slen+1);
    if(slen <= 0 || rownum == 1){ strncpy(res,s,slen);return res;}
    int onecircle = 2*rownum-2;
    int i;
    int resIndx = 0;
    int sIndx = 0;
    int skip = 0;

    for(i = 0; i < rownum; i++){
	skip = onecircle - 2*i;
        sIndx = i;
        res[resIndx++] = s[sIndx];

 	while((sIndx+=skip) < slen){
	    if(skip) res[resIndx++] = s[sIndx];
	    skip = onecircle - skip;
	}
    }
    res[resIndx] = '\0';
    return res;
}

