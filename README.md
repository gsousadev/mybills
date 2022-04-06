# My Bills

O MyBills é um sistema de gerenciamento de orçamento mensal.

Tem como base o modelo 50-15-35. Esse modelo de gestão esta bem explicado no link abaixo

https://blog.nubank.com.br/regra-50-15-35-financas/


Sendo assim teremos os seguintes regras de negocio a serem seguidas:

- Lançar a data do recebimento mensal
- Lançar o valor recebido no mes
- Lançar valores dentro das 3 categorias: 
    - 50% gastos essenciais (despesas fixas)
    - 15% prioridades financeiras (dividas ou investimentos)
    - 35% estilo de vida (despesas que podem ser cortadas)


O fluxo principal será o seguinte: 

    Usuário lança salario e data do mês
    
    Usuário lança um valor e coloca em uma das 3 categorias

    Sistema faz calculo e mostra em tela novos saldos para cada categoria

