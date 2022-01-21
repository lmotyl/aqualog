<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Tank extends Model
{
    use HasFactory;

    public const TABLE_NAME = 'tanks';

    public const ID = 'id';
    public const USER_ID = 'user_id';
    public const TANK_TYPE_ID = 'tank_type_id';
    public const NAME = 'name';

    protected $table = self::TABLE_NAME;

}
