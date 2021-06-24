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
    locale: 'en_US',
    is_org_admin: Faker::Boolean.boolean(true_ratio: 0.2),
    is_internal: Faker::Boolean.boolean(true_ratio: 0.5)
  }
end

File.open('./data/usersv1.json', 'w') do |f|
  f.write(JSON.pretty_generate(users))
end
