import React, { useState, useEffect } from 'react';
import { useForm, Controller } from 'react-hook-form';
import FormContainer from './FormContainer';
import GreenButton from '../Buttons/GreenButton';
import { FormProps, FieldConfig, Option } from './types';
import './Form.css';

const Form: React.FC<FormProps> = ({ title, fields, onSubmit, buttonText }) => {
  const { control, handleSubmit } = useForm();
  const [options, setOptions] = useState<{ [key: string]: Option[] }>({});

  if (buttonText === "") {
    buttonText = "Добавить"
  }

  useEffect(() => {
    const fetchOptionsForFields = async () => {
      const newOptions: { [key: string]: Option[] } = {};
      for (const field of fields) {
        if (field.fetchOptions) {
          newOptions[field.name] = await field.fetchOptions();
        }
      }
      setOptions(newOptions);
    };

    fetchOptionsForFields();
  }, [fields]);

  return (
    <FormContainer>
      <div>
        <h2 className="form-title">{title}</h2>
        <form onSubmit={handleSubmit(onSubmit)}>
          {fields.map((field) => (
            <div key={field.name} className="input-group">
                <Controller
                  name={field.name}
                  control={control}
                  defaultValue=""
                  render={({ field: controllerField }) => {
                    if (field.type === 'select') {
                      return (
                          <select {...controllerField} className="select-field">
                            <option defaultValue={""} value="" disabled hidden>{`${field.placeholder}`}</option>
                            {(options[field.name] || []).map(option => (
                              <option key={option.value} value={option.value}>
                                {option.label}
                              </option>
                            ))}
                          </select>
                        );
                      } else {
                          return (
                            <input
                              {...controllerField}
                              type={field.type}
                              placeholder={field.placeholder}
                              className="input-field"
                            />
                          );
                      }
                  }}
                />
            </div>
          ))}
          <GreenButton text={buttonText} icon={""} />
        </form>
      </div>
    </FormContainer>
  );
};

export default Form;