import React, { useState, useEffect } from 'react';
import { useForm, Controller } from 'react-hook-form';

interface FieldConfig {
  name: string;
  type: string;
  placeholder: string;
  fetchOptions?: () => Promise<{ value: string; label: string }[]>;
}

interface FormProps {
  title: string;
  fields: FieldConfig[];
  onSubmit: (data: any) => void;
}

const Form: React.FC<FormProps> = ({ title, fields, onSubmit }) => {
  const { control, handleSubmit } = useForm();

  // const [options, setOptions] = useState<{ [key: string]: { value: string; label: string }[] }>({});

  // useEffect(() => {
  //   const fetchOptionsForFields = async () => {
  //     const newOptions: { [key: string]: { value: string; label: string }[] } = {};
  //     for (const field of fields) {
  //       if (field.fetchOptions) {
  //         newOptions[field.name] = await field.fetchOptions();
  //       }
  //     }
  //     setOptions(newOptions);
  //   };

  //   fetchOptionsForFields();
  // }, [fields]);
  console.log("FIELDS", fields)
  
  return (
    <div>
      <h1>{title}</h1>
      <form onSubmit={handleSubmit(onSubmit)}>
        {fields.map((field) => (
          <div key={field.name} className="input-group">
            <Controller
              name={field.name}
              control={control}
              defaultValue=""
              render={({ field }) => {
                if (field.type === 'select') {
                  return (
                    <select {...field}>
                      <option value="">{`Select ${field.placeholder}`}</option>
                      {field.options.map(option => (
                        <option key={option.value} value={option.value}>
                          {option.label}
                        </option>
                      ))}
                    </select>
                  );
                } else {
                  console.log(field);
                  return (
                    <input
                      {...field}
                      type={field.type}
                      placeholder={field.placeholder}
                    />
                  );
                }
              }}
            />
          </div>
        ))}
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default Form;
