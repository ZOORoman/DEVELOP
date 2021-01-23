
from flask import Flask, render_template, request
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)                                       # Директива name
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///shop.db' # Обращаемся к SQLAlchemy
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False        # Для модификации настроек в db
db = SQLAlchemy(app)                                        # Присоединяем БД к объекту приложений

# БД (shop.db) - Таблицы - Записи
# Таблица:
#  id  title    price  isActive
#   1  Мишка    1500    True
#   2  Барашка  1000    False
#   3  Овечка   1000    True

class Item(db.Model): # Класс наследующий таблицу
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(100), nullable=False)
    price = db.Column(db.Integer, nullable=False)
    isActive = db.Column(db.Boolean, default=True)
    # text = db.Column(db.Text, default=False)

    def __repr__(self):
        return self.title
        # return f'Запись: {self.title}' - форматированный вывод записей

@app.route('/')                          # @app - итератор, / - переходим на главную
def index():                             # Метод который обрабатывает index
    items = Item.query.order_by(Item.price).all()
    return render_template('index.html', data=items) # Возвращает шаблон из папки templates

@app.route('/about')
def about():
    return render_template('about.html')

@app.route('/create', methods=['POST', 'GET'])
def create():
    if request.method == "POST":            # reqest - параметр котрорый отслеживает метод POST
        title = request.form['title']       # Получаем данные с метода POST
        price = request.form['price']

        item = Item(title=title, price=price)

        try:
            db.session.add(item)
            db.session.commit()
            return credits('/')
        except:
            return "Ошибка при добавлении"
    else:
        return render_template('create.html')

if __name__ == "__main__":
    app.run(debug=True)                 # Запуск сайта.
                                        # "debug=True" Когда стоит на сервере то лучше False,
                                        # так мы не покажем ошибки, которые возникают




