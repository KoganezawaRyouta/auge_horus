begin
  require 'standalone_migrations'
  StandaloneMigrations::Tasks.load_tasks
rescue LoadError => e
  puts "gem install standalone_migrations to get db:migrate:* tasks! (Error: #{e})"
end
