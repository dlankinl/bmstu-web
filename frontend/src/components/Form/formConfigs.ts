export const companyFields = [
  { name: 'name', type: 'text', placeholder: 'Название' },
  { name: 'city', type: 'text', placeholder: 'Город' },
  { name: 'activityFieldID', type: 'select', placeholder: 'Сфера деятельности',
    fetchOptions: async () => {
      
      // const response = await fetch('https://api.example.com/proficiency-levels');
      // const data = await response.json();
      // return data.map((item) => ({
      //   value: item.value,
      //   label: item.label,
      // }));
      return [{
        value: "testValue1",
        label: "testLabel1",
      }];
    },
  }
];

export const activityFieldFields = [
  { name: 'name', type: 'text', placeholder: 'Название' },
  { name: 'description', type: 'text', placeholder: 'Описание' },
  { name: 'cost', type: 'text', placeholder: 'Вес' },
];

export const signupFields = [
  { name: 'username', type: 'text', placeholder: 'Имя пользователя' },
  { name: 'password', type: 'password', placeholder: 'Пароль' },
  { name: 'verifyPassword', type: 'password', placeholder: 'Подтверждение пароля' },
]

export const signinFields = [
  { name: 'username', type: 'text', placeholder: 'Имя пользователя' },
  { name: 'password', type: 'password', placeholder: 'Пароль' },
]