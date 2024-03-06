import userEvent from '@testing-library/user-event';
import { render, screen } from '@testing-library/react';
import TextInput from './TextInput';

test('TextInput Component Test', async () => {
  const user = userEvent.setup();
  render(<TextInput />);

  const textElement = screen.getByText('Entered Text:');
  expect(textElement).toBeInTheDocument();

  const inputElement = screen.getByLabelText('Text Input');
  await user.type(inputElement, 'Hello World');
  expect(screen.getByText('Entered Text: Hello World')).toBeInTheDocument();
});
