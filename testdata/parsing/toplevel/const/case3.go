package main
const a int = 1; /* adfaf * sdf */
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  ConstDeclarationsImpl
    PsiElement(KEYWORD_CONST)('const')
    PsiWhiteSpace(' ')
    ConstSpecImpl
      LiteralIdentifierImpl
        PsiElement(IDENTIFIER)('a')
      PsiWhiteSpace(' ')
      TypeNameImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('int')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralIntegerImpl
          PsiElement(LITERAL_INT)('1')
  PsiElement(;)(';')
  PsiWhiteSpace(' ')
  PsiComment(ML_COMMENT)('/* adfaf * sdf */')
