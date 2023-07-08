import random
from datetime import datetime

def generate_random_string(length):
    letters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
    return ''.join(random.choice(letters) for _ in range(length))

def generate_insert_sql(num_records):
    sql_statements = []
    for _ in range(num_records):
        app_id = generate_random_string(8)
        spu_id = generate_random_string(8)
        spu_type = generate_random_string(10)
        goods_category_id = generate_random_string(8)
        goods_name = generate_random_string(15)
        goods_img = generate_random_string(12)
        price = random.randint(10, 100)
        price_line = price + random.randint(5, 20)
        goods_tag = generate_random_string(10)
        sale_status = random.randint(0, 1)
        sale_at = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        is_display = random.randint(0, 1)
        limit_purchase = random.randint(0, 10)
        stock_deduct_mode = random.randint(0, 1)
        is_deleted = random.randint(0, 1)
        created_at = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        updated_at = created_at

        sql = f"INSERT INTO `t_spu` (`app_id`, `spu_id`, `spu_type`, `goods_category_id`, `goods_name`, `goods_img`, `price`, `price_line`, `goods_tag`, `sale_status`, `sale_at`, `is_display`, `limit_purchase`, `stock_deduct_mode`, `is_deleted`, `created_at`, `updated_at`) VALUES ('{app_id}', '{spu_id}', '{spu_type}', '{goods_category_id}', '{goods_name}', '{goods_img}', {price}, {price_line}, '{goods_tag}', {sale_status}, '{sale_at}', {is_display}, {limit_purchase}, {stock_deduct_mode}, {is_deleted}, '{created_at}', '{updated_at}');"
        sql_statements.append(sql)

    return sql_statements

num_records = 30
insert_sql_statements = generate_insert_sql(num_records)

for sql_statement in insert_sql_statements:
    print(sql_statement)
