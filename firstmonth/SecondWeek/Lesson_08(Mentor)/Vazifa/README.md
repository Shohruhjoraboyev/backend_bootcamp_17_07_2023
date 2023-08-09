##Homeworks
1.Branches: 
fields:
    -id
    -name
    -address
methods:
    (1)create
    (2)update
    (3)getById
    (4)getAllda name, address bo'yicha search,pagination
    (5)delete
2.Staffs: 
fields:
    -id
    -branchId
    -tarifId
    -type (cashier,shop_assistant)
    -name
    -balance
methods:
     (1)create
     (2)update
     (3)getById
     (4)getAllda branchId,tariffId,type, name(search), balance(from,to) filter, pagination
     (5)delete
3. Staff_tariffs:
fields:
     - id
     - name
     - type (fixed, percent)
     - amoun_for_cash
     - amount_for_card
methods:
     (1)create
     (2)update
     (3)getById
     (4)getAllda  name(search), type filter, pagination
     (5)delete
4. Berilgan massiv palindrom yoki palindrom emasligini aniqlang.Palindrom deb o'ngdan chapga va chapdan o'nga bir xil o'qilishiga aytiladi
     Masalan: [1,2,1]=>true. [1,2,3,1] => false
5. Takrorlanmaydigan elementlardan iborat sliceda, low va high berilgan. low va high oralig'idagi arrayda yo'q bo'lgan sonlarni tartib bilan qaytaring