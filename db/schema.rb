# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20170514042844) do

  create_table "tickers", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.integer "last", null: false, unsigned: true
    t.integer "bid", null: false, unsigned: true
    t.integer "ask", null: false, unsigned: true
    t.integer "high", null: false, unsigned: true
    t.integer "low", null: false, unsigned: true
    t.float "volume", limit: 24, null: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "trades", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.integer "trade_id", null: false, unsigned: true
    t.float "amount", limit: 24, null: false
    t.bigint "rate", null: false
    t.string "order_type"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

end
