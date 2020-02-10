#!/usr/bin/env ruby

require 'faker'

users = 50.times.map do
  first_name = Faker::Name.first_name
  last_name = Faker::Name.last_name
  full_name = "#{first_name} #{last_name}"

  {
    id: Faker::Number.number(digits: 8),
    username: full_name.split(' ').join('.'),
    first_name: first_name,
    last_name: last_name,
    email: Faker::Internet.email(name: full_name, separators: '.', domain: 'example'),
    is_active: true,
    locale: 'en_US'
  }
end

File.open('./data/users.json', 'w') do |f|
  f.write(JSON.pretty_generate(users))
end
