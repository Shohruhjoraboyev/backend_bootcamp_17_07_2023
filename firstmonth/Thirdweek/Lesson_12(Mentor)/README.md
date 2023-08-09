##Homework:
1.Sales: 
fields:
    -id
    -branchId
    -shop_assistant_id
    -cashier_id
    -price
    -payment_type(card,cash)
    -status(success,cancel)
    -client_name
    -created_at
methods:
     (1)create
     (2)update
     (3)getById
     (4)getAllda branchId,shop_assistant_id,cashier_id,payment_type,status, client_name(search), price(from,to) filter, pagination
     (5)delete
2.Staff_transactions:
fields:
    -id
    -sale_id
    -staff_id
    -type(withdraw,topup)
    -source_type(sales,bonus)
    -amount
    -text
    -created_at
methods:
     (1)create
     (2)update
     (3)getById
     (4)getAllda sale_id,staff_id,type,source_type, created_at(from,to) filter, pagination
     (5)delete
3. Closure yordamida, men argumentga nechchi bersam shuni qo'shib qaytaradigan function yozing.
masalan, 
 adder(3) => 3
 adder(7) =>10
... ... ... ... ... ...
4. Berilgan n sonigacha barcha tub sonlarni chiqaring. Tub son - faqat o'ziga va 1ga bo'linadigan son(2,3,5,7.......)