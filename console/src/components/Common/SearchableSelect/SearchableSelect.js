import React from 'react';
import Select, { components, createFilter } from 'react-select';
import PropTypes from 'prop-types';

/*
 * Wrap the option generated by react-select and adds utility properties
 * */
const CustomOption = props => {
  return (
    <div
      title={props.data.description || ''}
      data-test={`data_test_column_type_value_${props.data.value}`}
    >
      <components.Option {...props} />
    </div>
  );
};

const getPrefixFilter = () => {
  const prefixFilterOptions = {
    matchFrom: 'start',
  };

  return createFilter(prefixFilterOptions);
};

/*
 * Searchable select box component
 *  1) options: Accepts options
 *  2) value: selectedValue
 *  3) onChange: function to call on change of value
 *  4) bsClass: Wrapper class
 *  5) customStyle: Custom style
 * */
const SearchableSelectBox = ({
  options,
  onChange,
  value,
  bsClass,
  styleOverrides,
  placeholder,
  filterOption,
}) => {
  /* Select element style customization */

  const customStyles = {};
  if (styleOverrides) {
    Object.keys(styleOverrides).forEach(comp => {
      customStyles[comp] = provided => {
        return {
          ...provided,
          ...styleOverrides[comp],
        };
      };
    });
  }

  let customFilter;
  switch (filterOption) {
    case 'prefix':
      customFilter = getPrefixFilter();
      break;
    default:
      customFilter = {};
  }

  return (
    <Select
      isSearchable
      components={{ Option: CustomOption }}
      classNamePrefix={`${bsClass}`}
      placeholder={placeholder}
      options={options}
      onChange={onChange}
      value={value}
      styles={customStyles}
      filterOption={customFilter}
    />
  );
};

SearchableSelectBox.propTypes = {
  value: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  options: PropTypes.array.isRequired,
  bsClass: PropTypes.string,
  customStyle: PropTypes.object,
  filterOption: PropTypes.object,
};

export default SearchableSelectBox;
