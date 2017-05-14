class TradesMigration < ActiveRecord::Migration[5.1]
  def change
    create_table :trades do |t|
      t.integer :trade_id, limit: 4, null: false, unsigned: true
      t.float :amount, null: false, null: false
      t.integer :rate, limit: 8, null: false
      t.string :order_type, limit: 255
      t.timestamps null: true
    end
  end
end
