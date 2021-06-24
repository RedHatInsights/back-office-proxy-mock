#!/usr/bin/env ruby

require 'faker'

first_name = Faker::Name.first_name
last_name = Faker::Name.last_name
full_name = "#{first_name} #{last_name}"
email = Faker::Internet.email(name: full_name, separators: '.', domain: 'example')
address_string = "\"#{full_name}\" #{email}"

users = {
  user: {
    id: Faker::Number.number(digits: 8),
    username: full_name.split(' ').join('.'),
    email: email,
    first_name: first_name,
    last_name: last_name,
    account_number: Faker::Number.number(digits: 8),
    address_string: address_string,
    is_active: true,
    is_org_admin: Faker::Boolean.boolean(true_ratio: 0.2),
    is_internal: Faker::Boolean.boolean(true_ratio: 0.5),
    locale: 'en_US',
    org_id: Faker::Number.number(digits: 8),
    display_name: full_name,
    type: 'basic'
  },
  mechanism: 'basic'
}

File.open('./data/auth.json', 'w') do |f|
  f.write(JSON.pretty_generate(users))
end
