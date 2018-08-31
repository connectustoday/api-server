import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IOpportunity extends Typegoose {

	@prop({ required: true })
        schema_version: number;

	@prop({ required: true })
	id: string;

	@prop({ required: true })
	organization: string;

	@prop({ required: true })
	name: string;

	@prop()
	description?: string;

	@prop()
	address?: IAddress;

	@prop({ required: true })
	is_signups_enabled: boolean;

	@prop()
	number_of_people_needed?: number;

	@prop()
	tags?: Array<string>;

	@prop()
	interested_users: Array<string>; //TODO

	@prop()
	shift_times: Array<[string, string]>;

	@prop()
	method_of_contact?: string;

	@prop({ required: true })
	created_at: number;
}
