export interface Option {
  value: string;
  label: string;
}

export interface FieldConfig {
  name: string;
  type: 'text' | 'select'; // Add other types as needed
  placeholder: string;
  options?: Option[]; // Only for select fields
  fetchOptions?: () => Promise<Option[]>; // Function to fetch options for select fields
}

export interface FormProps {
  title: string;
  fields: FieldConfig[];
  buttonText: string;
  onSubmit: (data: any) => void; // Adjust the type as necessary for your form data
}