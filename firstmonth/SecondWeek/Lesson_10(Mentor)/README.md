##Homework
1. O'tgan darsdagi hamma crudlarni cruds degan package ichida alohida filega ajratish
2. Hamma crudlarga created_at qo'shish, arrayga qo'shilgan vaqti yoziladi. va getAlllariga created_at bo'yicha descending orderda sort qo'shish ya'ni eng oxirgi qo'shilganlari birinchi chiqishi kerak. sort qilishni google qilamiz
3. Branches crudga founded_at(tashkil topgan vaqti),year(tashkil topganiga qancha bo'lgan) qo'shish va getById va getAllda yearni hisoblab qaytarish. Shunda:
Branch:
-id
-name
-address
-created_at
-founded_at
-year
create, updateda year ishlatilmidi faqat getlarda
4. Staffs crudga birthday_date(tug'ilgan kuni)ni, age(yosh)ni qo'shish va bunda ham getById va getAlllarda ageni hisoblab qaytarish
5. 
Clients:
-id
-name
-card(korzina)  Card

Card:
products   CardProducts

CardProducts:
-product_id
-size_id
-quantity

Product:
-id
-name
-sizes []Size

Size:
-id
-name(25sm,30sm,35sm)
-price

5.1. Har bitta klientni korzinas(card)ida qancha sumlik product borligini chiqaring
Masalan: 
Shohruh - 100 000
Islom      - 170 000
5.2. Qaysi product eng ko'p korzinkaga qo'shilgan, masalan:
Cola - 5ta
Fanta -2ta