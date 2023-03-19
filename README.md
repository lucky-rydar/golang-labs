### Система менеджменту гуртожитку

Система призначена для управління гуртожитком. Її основними ***цілями*** є надання можливості адміністрації гуртожитку керувати розселенням студентів у гортожитку:

- [x] мати інформацію про загруженість (кількість місць всього, вільні/зайняті місця); 
- [x] інформація про наповненість кімнати (доступність місця, людей, що проживають в кімнаті); 
- [x] мати відомості про мешканців гуртожитку (студентський квиток, угода, підписана мешканцем).

> Система призначена виключно для використання адміністрацією гуртожитку, мешканці не матимуть доступу до неї.

Система, в цілому, ділиться на два блоки: студент та його дані та дані про гуртожиток (кімнати, ліжко-місця).

![img](img/diagram.png)

#### Основні функції програми

> Що стосуються кімнат:
- [x] внесення в систему відомостей про наявні кімнати: її номер, чоловіча/жіноча, площа;
- [x] система обраховує оптимальну кількість місць в кімнаті та додає їх в систему як незайняті (розрахунок 4м2/чол);
- [x] можливість отримання даних про кімнату, зокрема, її номер, чоловіча/жіноча, площа, кількість вільних місць, мешканців;
- [x] можливість отримати дані про вільні місця (окремо жіночі/чоловічі).


> Що стосуються мешканців:
- [x] внесення в систему відомостей про студента: ім’я, прізвище, студентський;
- [x] присвоєння студенту вільного місця в гуртожитку;
- [x] внесення даних про угоду з мешканцем (дата початку і закінчення дії), можливість видалити прострочену угоду, якщо з мешканцем заключена інша.
